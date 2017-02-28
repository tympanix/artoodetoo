import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AdministrationComponent } from '../administration/administration.component';
import { DashboardComponent } from '../dashboard/dashboard.component';
import { StatisticsComponent } from '../statistics/statistics.component';

const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full'},
  { path: 'dashboard',      component: DashboardComponent},
  { path: 'administration', component: AdministrationComponent},
  { path: 'statistics',     component: StatisticsComponent}
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})

export class AppRoutingModule { }
