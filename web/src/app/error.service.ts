import { Injectable } from '@angular/core';
import { Notification } from './model';
import { Subject } from 'rxjs/Subject';

import { MdSnackBar } from '@angular/material'

@Injectable()
export class ErrorService {
  public errors: Subject<Notification> = new Subject<Notification>()

  constructor(private snackBar: MdSnackBar) { }

  push(notification: Notification) {
    this.errors.next(notification)
    this.snackBar.open(notification.error, "", {duration: 4000, extraClasses: ["snackbar-warning"]})
  }

}
