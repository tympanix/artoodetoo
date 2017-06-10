import { Component, OnInit, Input } from '@angular/core';

import { ApiService } from '../../api.service'
import { Event} from '../../model'

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
    if(event.checked == true){
      this.api.startEvent(this.event).subscribe(() => this.slideDisabled = false)

    } else{
      this.api.stopEvent(this.event).subscribe(() => this.slideDisabled = false)

    }
  }

}
