import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http'

import 'rxjs/add/operator/toPromise';

import { Unit } from './unit';
import { UNITS } from './mock-units';

@Injectable()
export class UnitService {

  constructor(private http: Http) {}

  getUnits(): Promise<Unit[]>{
    return this.http.get("/api/components")
                .toPromise()
                .then(this.extractData)
                .catch(this.handleError);
  }

  private extractData(res: Response) {
    let body = res.json();
    return body || { };
  }

  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error); // for demo purposes only
    return Promise.reject(error.message || error);
  }

  getMockUnits(): Promise<Unit[]> {
    return Promise.resolve(UNITS);
  }

  // See the "Take it slow" appendix
  getUnitsSlowly(): Promise<Unit[]> {
    return new Promise(resolve => {
      // Simulate server latency with 2 second delay
      setTimeout(() => resolve(this.getMockUnits()), 2000);
    });
  }
}
