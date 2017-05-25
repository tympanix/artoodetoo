import { Component, OnInit } from '@angular/core';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styles: []
})
export class LoginComponent implements OnInit {

  private password: string

  constructor(private api: ApiService) { }

  ngOnInit() {
  }

  login() {
    this.api.login(this.password).subscribe()
    console.log("LOGIN")
  }

}
