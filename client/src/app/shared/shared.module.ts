import { NgModule } from '@angular/core';

import { CustomSnackbarComponent } from './custom-snackbar/custom-snackbar.component';
import { MAT_SNACK_BAR_DEFAULT_OPTIONS } from '@angular/material';

import { MenuItems } from './menu-items/menu-items';
import {
  AccordionAnchorDirective,
  AccordionLinkDirective,
  AccordionDirective
} from './accordion';

import {
  AuthGuard, RolAdminGuard
} from '../guards/core';

import {
  LocalStorageService,
  SessionService,
  ErrorHandlerService,
  APIAccountService,
  APIDiskService,
  APIClientService,
  APIPDiscReservationService,
  FeedbackHandlerService,
  APIFreeInfoService,
} from '../services/core';
import { DemoMaterialModule } from '../demo-material-module';
import { Util } from '../models/util';
// import { APIAccountService } from '../services/api/account.services';


@NgModule({
  declarations: [
    AccordionAnchorDirective,
    AccordionLinkDirective,
    AccordionDirective,

    CustomSnackbarComponent,
  ],
  imports: [
    DemoMaterialModule,
  ],
  exports: [
    AccordionAnchorDirective,
    AccordionLinkDirective,
    AccordionDirective,
    CustomSnackbarComponent,
  ],
  entryComponents: [
    CustomSnackbarComponent,
  ],
  providers: [
    MenuItems,

    AuthGuard,
    RolAdminGuard,

    LocalStorageService,
    SessionService,
    ErrorHandlerService,
    FeedbackHandlerService,

    APIAccountService,
    APIDiskService,
    APIClientService,
    APIPDiscReservationService,
    APIFreeInfoService,

    Util,

    {provide: MAT_SNACK_BAR_DEFAULT_OPTIONS, useValue: {duration: 4000}}
  ]
})
export class SharedModule { }
