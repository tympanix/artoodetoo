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

import { ApiService } from './api.service';
import { TypeService } from './type.service'
import { ErrorService } from './error.service'
import { UnitComponent } from './unit/unit.component';

import { TaskResolver } from './resolvers/task-resolver.service';
import { IngredientComponent } from './ingredient/ingredient.component';
import { UnitDialog } from './dialogs/unitdialog/unitdialog.component';
import { TaskDialog } from './dialogs/taskdialog/taskdialog.component';
import { EventTemplateDialog } from './dialogs';
import { EventComponent } from './event/event.component';
import { EventIngredientComponent } from './event-ingredient/event-ingredient.component';

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
import { IngredientType } from './types/ingredient-type'

@NgModule({
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule,
    MaterialModule,
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
    EventIngredientComponent,
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
    BoolTypeComponent
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
  ],
  providers: [ApiService, TaskResolver, ErrorService, TypeService],
  bootstrap: [AppComponent]
})
export class AppModule { }
