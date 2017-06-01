import { Component, OnInit } from '@angular/core';
import { TypesComponent } from '../types.component'
import { FacebookService, InitParams, LoginResponse, LoginOptions, LoginStatus } from 'ngx-facebook'

@Component({
  templateUrl: './facebook-token.component.html',
  styles: []
})
export class FacebookTokenComponent extends TypesComponent {
  facebookToken: string
  loginStatus: string
  options: LoginOptions

  constructor(private fb: FacebookService) {
    super()
  }

  loginWithFacebook(): void {

    // login with options
    this.fb.login(this.options)
      .then((response: LoginResponse) => {
        this.loginStatus = response.status
        if(response.status == 'connected'){

          this.ingredient.value = response.authResponse.accessToken
        } else{
          return new Error();
        }
      })
      .catch((error: any) => console.error(error))

  }

  logoutWithFacebook(): void {
    this.fb.logout()
      .then(() => {
        this.loginStatus = "disconnected"
        this.ingredient.value = null
      })
  }

}
