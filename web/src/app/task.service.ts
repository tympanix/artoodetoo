import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http'

import 'rxjs/add/operator/toPromise';

import { Unit } from './unit';
import { Task } from './task';
import { TASK } from './mock/task';

@Injectable()
export class TaskService {

  constructor(private http: Http) {}

  createTask(task: Task): Promise<Task>{
    let headers = new Headers({ 'Content-Type': 'application/json' });
    let options = new RequestOptions({ headers: headers });

    return this.http.post("api/tasks", { task }, options)
                .toPromise()
                .then(this.extractData)
                .catch(this.handleError);
  }

  createMockTask(): Promise<Task>{
    let headers = new Headers({ 'Content-Type': 'application/json' });
    let options = new RequestOptions({ headers: headers });
    console.log(TASK);
    return this.http.post("api/tasks", TASK, options)
                .toPromise()
                .then(this.extractData)
                .catch(this.handleError);
  }

  private extractData(res: Response) {
    let body = res.status;
    return body || { };
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

}
