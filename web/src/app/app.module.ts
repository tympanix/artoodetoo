import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';
import { MaterialModule } from '@angular/material'

import { AppRoutingModule } from './app-routing/app-routing.module';

import { AppComponent } from './app.component';
import { SiteHeaderComponent } from './site-header/site-header.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { AdministrationComponent } from './administration/administration.component';
import { StatisticsComponent } from './statistics/statistics.component';

import { ApiService }           from './api.service';
import { UnitComponent } from './unit/unit.component';

import { TaskResolver } from './resolvers/task-resolver.service';
import { DropdownComponent } from './dropdown/dropdown.component';

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule,
    MaterialModule
  ],
  declarations: [
    AppComponent,
    SiteHeaderComponent,
    DashboardComponent,
    AdministrationComponent,
    StatisticsComponent,
    UnitComponent,
    DropdownComponent
  ],
  providers: [ApiService, TaskResolver],
  bootstrap: [AppComponent]
})
export class AppModule { }
