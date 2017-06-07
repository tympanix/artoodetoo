import { Injectable } from '@angular/core';
import { ApiService } from './api.service';
import { Log } from './model';

@Injectable()
export class LogService {
  errors: number = 0
  logs: Log[] = []

  constructor(private api: ApiService) {
    this.api.logs.subscribe(l => this.logs.push(l))
    this.api.logs.filter(l => l.isError()).subscribe(() => this.errors++)
  }

  clear() {
    this.logs = []
    this.errors = 0
    this.api.clearLog().subscribe()
  }

}
