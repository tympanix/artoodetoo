import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http'

import { ReplaySubject } from 'rxjs/ReplaySubject';
import { Observable } from 'rxjs/Observable';
import { MdSnackBar } from '@angular/material';

import 'rxjs/add/operator/map'

import { Task, Unit, Data, Event } from './model';

@Injectable()
export class ApiService {

  public tasks: ReplaySubject<Task[]> = new ReplaySubject<Task[]>(1)
  public units: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)
  public templateEvents: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)
  public events: ReplaySubject<Event[]> = new ReplaySubject<Event[]>(1)

  private options: RequestOptions

  constructor(private http: Http, private snackBar: MdSnackBar) {
    let headers = new Headers({ 'Content-Type': 'application/json' });
    this.options = new RequestOptions({ headers: headers });

    // this.getTasks()
    // this.getUnits()
    // this.getTemplateEvents()
    // this.getEvents()

    this.getAll()
  }

  private extractData<T>(self: this): (res: Response) => T {
    return function(res: Response): T {
      if (res.status != 200) {
        //self.snackBar.open("Error", res.text(), {duration: 4000})
        throw new Error(res.text())
      }
      let body:T = res.json();
      if (body) {
        return body
      } else {
        throw new Error(res.text())
      }
    }
  }

  private handleError(self: this) {
    return function(error: any): Promise<any> {
      console.error('An error occurred', error); // for demo purposes only
      self.snackBar.open("Error", error.message || error, {duration: 4000})
      return Promise.reject(error.message || error);
    }
  }

  createTask(task: Task): Observable<boolean> {
    return this.http.post("api/tasks", task.toJson(), this.options)
      .map(this.extractData<string>(this))
      .do(() => this.getAll())
      .do(() => {
          this.snackBar.open(task.name + " has been created!", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      })
      .do(() => task.isSaved = true)
      .catch(this.handleError(this))
  }

  getTasks(): Observable<Task[]> {
    this.http.get("/api/tasks")
      .map(this.extractData<Task[]>(this))
      .map(json => json.map(data => Task.fromJson(data)))
      .catch(this.handleError(this))
      .subscribe(tasks => this.tasks.next(tasks));
    return this.tasks
  }

  getTemplateEvents(): Observable<Unit[]>{
    this.http.get("/api/all_events")
      .map(this.extractData<Unit[]>(this))
      .map(json => json.map(data => Unit.fromJson(data)))
      .catch(this.handleError(this))
      .subscribe(events => this.templateEvents.next(events));
    return this.templateEvents
  }

  getAll() {
    return this.http.get("/api/all")
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
    this.http.get("/api/events")
      .map(this.extractData<Unit[]>(this))
      .map(json => json.map(data => Unit.fromJson(data)))
      .catch(this.handleError(this))
      .subscribe(events => this.events.next(events));
      console.log("event")
    return this.events
  }

  runTask(task: Task): Observable<boolean> {
    return this.http.post("/api/tasks/" + task.name, {})
      .map(res => res.ok)
      .do(bool => {
          this.snackBar.open(task.name + " has been deployed!", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      })
      .catch(this.handleError(this))
  }

  stopTask(task: Task){
    // To be filled
  }

  updateTask(task: Task): Observable<boolean> {
    return this.http.put("/api/tasks", task.toJson(), this.options)
      .map(res => res.ok)
      .do(bool => {
          this.snackBar.open(task.name + " has been updated!", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      })
      .catch(this.handleError(this))
  }

  getUnits(): Observable<Unit[]> {
    this.http.get("/api/units")
      .map(this.extractData<Unit[]>(this))
      .map(json => json.map(data => Unit.fromJson(data)))
      .catch(this.handleError(this))
      .subscribe(units => this.units.next(units))
    return this.units
  }

  deleteTask(task: Task): Observable<boolean> {
    return this.http.delete("api/tasks/" + task.name, {})
      .map(res => res.ok)
      .do(bool => {
          this.snackBar.open(task.name + " has been deleted!", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      })
      .catch(this.handleError(this))
  }

  saveEvent(event: Unit): Observable<boolean>{
    return this.http.post("api/events", event.toJson(), this.options)
      .map(res => res.ok)
      .do(bool => {
          this.snackBar.open(event.name + " has been created!", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      })
      .catch(this.handleError(this))
  }

}
