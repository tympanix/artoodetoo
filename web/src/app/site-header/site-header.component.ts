import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service'

@Component({
  selector: 'app-site-header',
  templateUrl: './site-header.component.html',
  styles: []
})
export class SiteHeaderComponent implements OnInit {
  errors: number = 0

  constructor(private api: ApiService) {
    let self = this
    this.api.logs.filter(l => l.type == "error").subscribe(l => self.errors++)
  }

  ngOnInit() {
  }

  logout() {
    this.api.logout()
  }

}
