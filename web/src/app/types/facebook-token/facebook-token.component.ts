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

    // Hardcode to static type
    let initParams: InitParams = {
      appId: '1890556247880818',
      xfbml: true,
      version: 'v2.9'
    }

    fb.init(initParams)

    this.options = {
      scope: 'public_profile,user_friends,email,pages_show_list',
      return_scopes: true,
      enable_profile_selector: true
    };
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
