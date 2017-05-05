import { Component, OnInit, Input } from '@angular/core';
import { Ingredient, Input as UnitInput } from '../model';
import { FacebookService, InitParams, LoginResponse, LoginOptions, LoginStatus } from 'ngx-facebook'

@Component({
  selector: 'event-ingredient',
  templateUrl: './event-ingredient.component.html',
  styles: []
})
export class EventIngredientComponent implements OnInit {
  @Input() input: UnitInput
  @Input() ingr: Ingredient

  // Cron variables
  timeType: number = 0
  minute: number[]
  hour: number[]
  selectedType: number
  selectedNumber: number

  // Facebook variables
  facebookToken: string
  loginStatus: string

  constructor(private fb: FacebookService) {
    this.minute = [0,5,10,15,20,25,30,35,40,45,50,55]
    this.hour = [0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23]

    // Hardcore to static type
    let initParams: InitParams = {
      appId: '1890556247880818',
      xfbml: true,
      version: 'v2.9'
    }

    fb.init(initParams)
  }

  ngOnInit() {
    this.fb.getLoginStatus()
      .then((status:LoginStatus) => (this.loginStatus = status.status, this.ingr.value = status.authResponse.accessToken))
  }

  unitToNumber() {
    this.selectedType = +this.selectedType
  }

  typeToNumber(){
    this.timeType = +this.timeType
  }

  // Cron functions
  updateIngrValue(){
    this.ingr.value = "@every " + this.selectedNumber + (this.selectedType == 0 ? "m" : "h")
  }

  loginWithFacebook(): void {

    this.fb.login()
      .then((response: LoginResponse) => this.ingr.value = response.authResponse.accessToken)
      .catch((error: any) => console.error(error));

  }

}
