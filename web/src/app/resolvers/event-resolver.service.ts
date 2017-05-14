import { Injectable } from '@angular/core';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';

import { ReplaySubject } from 'rxjs/ReplaySubject';

import { ApiService } from '../api.service';
import { Task, Event } from '../model';

@Injectable()
export class EventResolver implements Resolve<Event> {

  constructor(private api: ApiService) { }

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    let obs: ReplaySubject<Event> = new ReplaySubject<Event>(1)
    let id = route.params['event']
    this.api.events.subscribe((events) => {
      obs.next(this.eventById(id, events))
      obs.complete()
    })
    return obs
  }

  eventById(id: string, events: Event[]): Event {
    return events.find((event) => event.uuid == id)
  }

}
