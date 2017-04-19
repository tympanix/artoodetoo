import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http'

import { ReplaySubject } from 'rxjs/ReplaySubject';
import { Observable } from 'rxjs/Observable';
import { MdSnackBar } from '@angular/material';

import 'rxjs/add/operator/map'

import { Task, Unit } from './model';

@Injectable()
export class ApiService {

  public tasks: ReplaySubject<Task[]> = new ReplaySubject<Task[]>(1)
  public units: ReplaySubject<Unit[]> = new ReplaySubject<Unit[]>(1)

  private options: RequestOptions

  constructor(private http: Http, private snackBar: MdSnackBar) {
    let headers = new Headers({ 'Content-Type': 'application/json' });
    this.options = new RequestOptions({ headers: headers });

    this.getTasks()
    this.getUnits()
  }

  private extractData<T>(): (res: Response) => T {
    return function(res: Response): T {
      let body:T = res.json();
      return body || {} as T;
    }
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    console.log(this);
    this.snackBar.open("Error", error.message || error, {duration: 4000})
    return Promise.reject(error.message || error);
  }

  createTask(task: Task): Observable<boolean> {
    return this.http.post("api/tasks", task.toJson(), this.options)
      .map(res => res.ok)
      .do(bool => {
          this.snackBar.open(task.name + " has been created!", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      })
      .catch(this.handleError)
  }

  getTasks(): Observable<Task[]> {
    this.http.get("/api/tasks")
      .map(this.extractData<Task[]>())
      .map(json => json.map(data => Task.fromJson(data)))
      .catch(this.handleError)
      .subscribe(tasks => this.tasks.next(tasks));
    return this.tasks
  }

  runTask(task: Task): Observable<boolean> {
    return this.http.post("/api/tasks/" + task.name, {})
      .map(res => res.ok)
      .do(bool => {
          this.snackBar.open(task.name + " has been deployed!", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      })
      .catch(this.handleError)
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
      .catch(this.handleError)
  }

  getUnits(): Observable<Unit[]> {
    this.http.get("/api/units")
      .map(this.extractData<Unit[]>())
      .map(json => json.map(data => Unit.fromJson(data)))
      .catch(this.handleError)
      .subscribe(units => this.units.next(units))
    return this.units
  }

  deleteTask(task: Task): Observable<boolean> {
    return this.http.delete("api/tasks/" + task.name, {})
      .map(res => res.ok)
      .do(bool => {
          this.snackBar.open(task.name + " has been deleted!", "", {duration: 4000, extraClasses: ["snackbar-success"]})
      })
      .catch(this.handleError)
  }

}
