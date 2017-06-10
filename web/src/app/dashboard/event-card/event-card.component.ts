import { Component, OnInit, Input } from '@angular/core';

import { Event} from '../../model'

@Component({
  selector: 'app-event-card',
  templateUrl: './event-card.component.html',
  styles: []
})
export class EventCardComponent implements OnInit {
  @Input() event: Event
  slideDisabled: boolean = false;

  constructor() { }

  ngOnInit() {
  }

  changeStatus(event){
    this.slideDisabled = true;
    if(event.checked == true){
      console.log("enable")

      // put this in callback
      this.slideDisabled = false;
    } else{
      console.log("disable")

      //put this in callback
      this.slideDisabled = false;
    }
  }

}
