import { Component, OnInit, NgZone } from '@angular/core';
import { TypesComponent } from '../types.component'

declare var gapi: any

@Component({
  templateUrl: './google-token.component.html',
  styles: []
})
export class GoogleTokenComponent extends TypesComponent {
  profile: any
  auth: any
  loggedIn: boolean

  constructor(private _zone: NgZone) {
    super()
    this.handleClientLoad()
  }

  handleClientLoad() {
    // Loads the client library and the auth2 library together for efficiency.
    // Loading the auth2 library is optional here since `gapi.client.init` function will load
    // it if not already loaded. Loading it upfront can save one network request.
    gapi.load('client:auth2', this.initClient.bind(this));
  }

  initClient() {
    let self: this = this
    // Initialize the client with API key and People API, and initialize OAuth with an
    // OAuth 2.0 client ID and scopes (space delimited string) to request access.
    gapi.client.init({
      discoveryDocs: ["https://www.googleapis.com/discovery/v1/apis/drive/v3/rest"],
      clientId: '756274658809-gmj1eu581con3av5vae7v6as8ndd9o7l.apps.googleusercontent.com',
      scope: 'https://www.googleapis.com/auth/drive'
    }).then(function() {
      // Listen for sign-in state changes.
      gapi.auth2.getAuthInstance().isSignedIn.listen(self.updateSigninStatus.bind(self));

      // Handle the initial sign-in state.
      self.updateSigninStatus(gapi.auth2.getAuthInstance().isSignedIn.get());
    });
  }

  updateSigninStatus(isSignedIn) {
    // When signin status changes, this function is called.
    // If the signin status is changed to signedIn, we make an API call.
    if (isSignedIn) {
      this.update()
    } else {
      this.remove()
    }
  }

  update() {
    let user = gapi.auth2.getAuthInstance().currentUser.get()

    if (user.isSignedIn()) {
      this._zone.run(() => {
        this.loggedIn = true
        this.profile = user.getBasicProfile()
        this.auth = user.getAuthResponse()
        this.ingredient.value = this.auth.access_token
      })
    }
  }

  remove() {
    this._zone.run(() => {
      this.loggedIn = false
      this.profile = undefined
      this.auth = undefined
    })
  }

  handleSignInClick(event) {
    // Ideally the button should only show up after gapi.client.init finishes, so that this
    // handler won't be called before OAuth is initialized.
    gapi.auth2.getAuthInstance().signIn();
  }

  handleSignOutClick(event) {
    gapi.auth2.getAuthInstance().signOut();
  }

}
