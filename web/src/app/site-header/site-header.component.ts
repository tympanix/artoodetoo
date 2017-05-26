import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service'

@Component({
  selector: 'app-site-header',
  templateUrl: './site-header.component.html',
  styles: []
})
export class SiteHeaderComponent implements OnInit {

  constructor(private api: ApiService) { }

  ngOnInit() {
  }

  logout() {
    this.api.logout()
  }

}
