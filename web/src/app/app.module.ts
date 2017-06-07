import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';
import { MaterialModule } from '@angular/material'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing/app-routing.module';

import { AppComponent } from './app.component';
import { SiteHeaderComponent } from './site-header/site-header.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { AdministrationComponent } from './administration/administration.component';

import { ApiService } from './api.service';
import { TypeService } from './type.service'
import { ErrorService } from './error.service'
import { DebuggerService } from './debugger.service'
import { UnitComponent } from './unit/unit.component';

import { TaskResolver } from './resolvers/task-resolver.service';
import { EventResolver } from './resolvers/event-resolver.service';
import { IngredientComponent } from './ingredient/ingredient.component';
import { UnitDialog } from './dialogs/unitdialog/unitdialog.component';
import { TaskDialog } from './dialogs/taskdialog/taskdialog.component';
import { EventTemplateDialog } from './dialogs';
import { EventComponent } from './event/event.component';

import { FacebookModule } from 'ngx-facebook';
import { TaskeditorComponent } from './editor/taskeditor/taskeditor.component';
import { EventeditorComponent } from './editor/eventeditor/eventeditor.component';
import { EventDialog } from './dialogs/eventdialog/eventdialog.component';
import { OptionDialog } from './dialogs/optiondialog/optiondialog.component';
import { ErrorDialog } from './dialogs/errordialog/errordialog.component';
import { CycleDialog } from './dialogs/cycledialog/cycledialog.component';
import { LoginComponent } from './login/login.component';
import { IntegerTypeComponent } from './types/integer-type/integer-type.component';
import { TypeDirective } from './types/type.directive';
import { StringTypeComponent } from './types/string-type/string-type.component';
import { TypesComponent } from './types/types.component';
import { BoolTypeComponent } from './types/bool-type/bool-type.component'
import { IngredientType } from './types/ingredient-type';
import { CronTimeComponent } from './types/cron-time/cron-time.component';
import { FacebookTokenComponent } from './types/facebook-token/facebook-token.component';
import { TypeeditorComponent } from './editor/typeeditor/typeeditor.component';
import { FloatTypeComponent } from './types/float-type/float-type.component';
import { EventDashboardComponent } from './dashboard/event-dashboard/event-dashboard.component';
import { GoogleTokenComponent } from './types/google-token/google-token.component';
import { DebuggerComponent } from './debugger/debugger.component'

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule,
    MaterialModule,
    BrowserAnimationsModule,
    FacebookModule.forRoot()
  ],
  declarations: [
    AppComponent,
    SiteHeaderComponent,
    DashboardComponent,
    AdministrationComponent,
    UnitComponent,
    IngredientComponent,
    UnitDialog,
    TaskDialog,
    EventTemplateDialog,
    EventComponent,
    TaskeditorComponent,
    EventeditorComponent,
    EventDialog,
    OptionDialog,
    ErrorDialog,
    CycleDialog,
    LoginComponent,
    IntegerTypeComponent,
    TypeDirective,
    StringTypeComponent,
    BoolTypeComponent,
    CronTimeComponent,
    FacebookTokenComponent,
    TypeeditorComponent,
    FloatTypeComponent,
    EventDashboardComponent,
    GoogleTokenComponent,
    DebuggerComponent
  ],
  entryComponents: [
    UnitDialog,
    TaskDialog,
    EventTemplateDialog,
    EventDialog,
    OptionDialog,
    ErrorDialog,
    CycleDialog,
    IntegerTypeComponent,
    StringTypeComponent,
    BoolTypeComponent,
    FacebookTokenComponent,
    CronTimeComponent,
    FloatTypeComponent,
    GoogleTokenComponent
  ],
  providers: [ApiService, TaskResolver, ErrorService, TypeService, EventResolver],
  bootstrap: [AppComponent]
})
export class AppModule { }
