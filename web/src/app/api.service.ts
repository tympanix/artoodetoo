import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Http, Response, Headers, RequestOptions } from '@angular/http'

import { ReplaySubject } from 'rxjs/ReplaySubject';
import { Observable } from 'rxjs/Observable';
import { MdSnackBar, MdDialog, MdSnackBarRef, SimpleSnackBar } from '@angular/material';

import 'rxjs/add/operator/map'
import 'rxjs/add/operator/mergeMap'

import { ErrorDialog } from './dialogs/errordialog/errordialog.component'
import { Task, Unit, Data, Event, Log } from './model';

@Injectable()
export class ApiService {
  private lastLog: number = 0

  public tasks: ReplaySubject<Task[]> = new ReplaySubject<Task[]>(1)
  public units: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)
  public templateEvents: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)
  public events: ReplaySubject<Event[]> = new ReplaySubject<Event[]>(1)
  public logs: ReplaySubject<Log> = new ReplaySubject<Log>(256)

  private options: RequestOptions
  private token: string

  constructor(private http: Http, private router: Router, private snackBar: MdSnackBar, public dialog: MdDialog) {
    this.token = localStorage.getItem("Token")
    this.createHeaders()
    this.getAll()
    this.updateLogs()

    setInterval(() => {
      this.updateLogs()
    }, 5000)
  }

  private createHeaders() {
    let headers = new Headers({
      'Content-Type': 'application/json',
      'Authentication': this.token
    });
    this.options = new RequestOptions({ headers: headers });
  }

  private setToken(token: string) {
    this.token = token
    localStorage.setItem("Token", token)
    this.createHeaders()
  }

  private extractData<T>(self: this): (res: Response) => T {
    return function(res: Response): T {
      self.checkSuccess(self)(res)
      let body:T = res.json();
      if (body) {
        return body
      } else {
        throw new Error(res.text())
      }
    }
  }

  private checkSuccess(self: this): (r: Response) => Response {
    return function(res: Response): Response {
      if (res.status != 200) {
        throw new Error(res.text())
      }
      return res
    }
  }

  private openErrorDialog(self: this, res: Response) {
    return function() {
      let dialogRef = self.dialog.open(ErrorDialog, {
        width: '750px',
        data: res
      });
      dialogRef.afterClosed().subscribe();
    }
  }

  private updateLogs() {
    this.getLog().subscribe(logs => logs.forEach(l => this.logs.next(l)))
  }

  private success(self: this, message: string): MdSnackBarRef<SimpleSnackBar> {
    return self.snackBar.open(message, "", {duration: 4000, extraClasses: ["snackbar-success"]})
  }

  private error(self: this, message: string): MdSnackBarRef<SimpleSnackBar> {
    return self.snackBar.open(message, "", {duration: 4000, extraClasses: ["snackbar-error"]})
  }

  private handleError(self: this) {
    return function(error: any): Promise<any> {
      console.error('An error occurred', error); // for demo purposes only
      var message: String
      if (error instanceof String) {
        message = error
      } else if (error instanceof Response) {
        if (error.status == 401) {
          self.router.navigateByUrl("/login")
        }
        message = error.text()
      }
      var snakcRef = self.snackBar.open(message as string, "Debug", {duration: 4000, extraClasses: ["snackbar-error"]})
      snakcRef.onAction().subscribe(self.openErrorDialog(self, error));
      return Promise.reject(error.message || error);
    }
  }

  private getLog() {
    let self = this
    return this.http.get("api/logs?t="+this.lastLog, this.options)
      .map(this.extractData<Log[]>(this))
      .map(logs => logs.map(log => Log.fromJson(log)))
      .do(logs => self.updateLastLog.bind(self)(logs[logs.length - 1]))
  }

  private updateLastLog(log: Log) {
    if (log) {
      this.lastLog = log.time
    }
  }

  clearLog() {
    return this.http.delete("api/logs", this.options)
      .map(this.checkSuccess(this))
  }

  login(username: string, password: string): Observable<string> {
    return this.http.post("api/login", {"username": username, "password": password}, this.options)
      .map(this.checkSuccess(this))
      .map((resp: Response) => resp.text())
      .do((token: string) => this.setToken(token))
      .do(() => this.success(this, "Login successful" ))
  }

  logout() {
    localStorage.removeItem("Token")
    this.token = ""
    this.createHeaders()
    this.router.navigateByUrl("/login")
  }

  createTask(task: Task): Observable<boolean> {
    return this.http.post("api/tasks", task.toJson(), this.options)
      .map(this.checkSuccess(this))
      .do(() => this.getTasks())
      .do(() => this.success(this, task.name + " has been created!"))
      .do(() => task.isSaved = true)
      .catch(this.handleError(this))
  }

  getTasks(): Observable<Task[]> {
    let _tasks: Task[]
    this.http.get("/api/tasks", this.options)
      .map(this.extractData<Task[]>(this))
      .map(json => json.map(data => Task.fromJson(data)))
      .do(tasks => _tasks = tasks)
      .mergeMap(tasks => this.events)
      .do(e => _tasks.forEach(t => t.resolveEvent(e)))
      .catch(this.handleError(this))
      .subscribe(() => this.tasks.next(_tasks))
    return this.tasks
  }

  getTemplateEvents(): Observable<Unit[]>{
    this.http.get("/api/all_events", this.options)
      .map(this.extractData<Unit[]>(this))
      .map(json => json.map(data => Event.fromJson(data)))
      .catch(this.handleError(this))
      .subscribe(events => this.templateEvents.next(events));
    return this.templateEvents
  }

  getAll() {
    return this.http.get("/api/all", this.options)
      .map(this.extractData<Data>(this))
      .map(json => Data.fromJson(json))
      .catch(this.handleError(this))
      .subscribe(data => {
        this.events.next(data.events)
        this.tasks.next(data.tasks)
        this.units.next(data.actions)
        this.templateEvents.next(data.eventtemplates)
      })
  }

  getEvents(): Observable<Unit[]>{
    this.http.get("/api/events", this.options)
      .map(this.extractData<Unit[]>(this))
      .map(json => json.map(data => Event.fromJson(data)))
      .catch(this.handleError(this))
      .subscribe(events => this.events.next(events));
    return this.events
  }

  updateTask(task: Task): Observable<boolean> {
    return this.http.put("/api/tasks", task.toJson(), this.options)
      .map(res => res.ok)
      .do(bool => this.success(this, task.name + " has been updated!"))
      .catch(this.handleError(this))
  }

  getUnits(): Observable<Unit[]> {
    this.http.get("/api/units", this.options)
      .map(this.extractData<Unit[]>(this))
      .map(json => json.map(data => Unit.fromJson(data)))
      .catch(this.handleError(this))
      .subscribe(units => this.units.next(units))
    return this.units
  }

  deleteTask(task: Task): Observable<boolean> {
    return this.http.delete("api/tasks/" + task.uuid, this.options)
      .map(this.checkSuccess(this))
      .do(bool => this.success(this, task.name + " has been deleted!"))
      .do(() => this.getAll())
      .catch(this.handleError(this))
  }

  saveEvent(event: Unit): Observable<boolean>{
    return this.http.post("api/events", event.toJson(), this.options)
      .map(res => res.ok)
      .do(() => this.getEvents())
      .do(bool => this.success(this, event.name + " has been created!"))
      .catch(this.handleError(this))
  }

  startEvent(event: Event): Observable<boolean>{
    return this.http.post("api/events/" + event.uuid + "/start", null, this.options)
      .map(res => res.ok)
      .do(() => this.getEvents())
      .do(bool => this.success(this, event.name + " has been enabled!"))
      .catch(this.handleError(this))
  }

  stopEvent(event: Event): Observable<boolean>{
    return this.http.post("api/events/" + event.uuid + "/stop", null, this.options)
      .map(res => res.ok)
      .do(() => this.getEvents())
      .do(bool => this.success(this, event.name + " has been disabled!"))
      .catch(this.handleError(this))
  }

}
