import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styles: []
})
export class LoginComponent implements OnInit {

  private password: string
  private error: string

  constructor(private api: ApiService, private router: Router) { }

  ngOnInit() {
  }

  private loginSuccess() {
    this.api.getAll()
    this.router.navigateByUrl('/')
  }

  private loginError() {
    this.error = "Invalid credentials"
  }

  login() {
    let self = this
    this.api.login(this.password).subscribe(
      () => self.loginSuccess(),
      () => self.loginError()
    )
  }

}
