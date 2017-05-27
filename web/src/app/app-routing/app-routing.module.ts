import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AdministrationComponent } from '../administration/administration.component';
import { DashboardComponent } from '../dashboard/dashboard.component';
import { LoginComponent } from '../login/login.component';

import { TaskResolver } from '../resolvers/task-resolver.service';
import { EventResolver } from '../resolvers/event-resolver.service';


const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full'},
  { path: 'dashboard', component: DashboardComponent},
  { path: 'login', component: LoginComponent},
  { path: 'administration', component: AdministrationComponent},
  { path: 'administration/task', component: AdministrationComponent},
  { path: 'administration/event', component: AdministrationComponent},
  { path: 'administration/task/:task', component: AdministrationComponent, resolve: { task: TaskResolver }},
  { path: 'administration/event/:event', component: AdministrationComponent, resolve: { task: EventResolver }},
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})

export class AppRoutingModule { }
