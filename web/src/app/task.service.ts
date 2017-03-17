import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http'
import {MdSnackBar} from '@angular/material';

import 'rxjs/add/operator/toPromise';
import 'rxjs/add/operator/map'

import { Task, Unit } from './task';
import { TASK } from './mock/task';

@Injectable()
export class TaskService {

  constructor(private http: Http, private snackBar: MdSnackBar) {}

  createTask(task: Task): Promise<boolean>{
    let headers = new Headers({ 'Content-Type': 'application/json' });
    let options = new RequestOptions({ headers: headers });

    return this.http.post("api/tasks", { task }, options)
                .toPromise()
                .then((res:Response) => res.ok)
                .catch(this.handleError);
  }

  getTasks(): Promise<Task[]>{
    return this.http.get("/api/tasks")
                .toPromise()
                .then(this.extractData)
                .catch(this.handleError);
  }

  runTask(task: Task): Promise<boolean> {
      console.log("Running task", task.name)
      return this.http.post("/api/tasks/"+task.name, {})
                .toPromise()
                .then(res => res.ok)
                .then(() => this.snackBar.open("Your task has been queued! Depending on the queue it may take a while..","", {duration: 4000}))
                .then(() => true)
                .catch(this.handleError)
  }

  private extractData(res: Response) {
    let body = res.json();
    return body || { };
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  createMockTask(): Promise<Task>{
    let headers = new Headers({ 'Content-Type': 'application/json' });
    let options = new RequestOptions({ headers: headers });
    return this.http.post("api/tasks", TASK, options)
                .toPromise()
                .then(this.extractData)
                .catch(this.handleError);
  }
}
