import { Component, OnInit, Input } from '@angular/core';

import { Observable } from 'rxjs/Observable';
import { ApiService } from '../../api.service'
import { Event } from '../../model'

@Component({
  selector: 'app-event-card',
  templateUrl: './event-card.component.html',
  styles: []
})
export class EventCardComponent implements OnInit {
  @Input() event: Event
  slideDisabled: boolean = false;

  constructor(private api: ApiService) { }

  ngOnInit() {
  }

  changeStatus(event){
    this.slideDisabled = true;
    let toggle: Observable<boolean>
    if(event.checked == true){
      toggle = this.api.startEvent(this.event)
    } else{
      toggle = this.api.stopEvent(this.event)
    }
    toggle.finally(() => this.slideDisabled = false).subscribe()
  }

}
