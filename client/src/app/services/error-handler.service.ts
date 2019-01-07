import { Injectable } from '@angular/core';
import { ResponseType } from '@angular/http';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material';
import { Router } from '@angular/router';

import { SessionService } from './session.service';
import { CustomSnackbarComponent } from '../shared/custom-snackbar/custom-snackbar.component';
import { isNullOrUndefined } from 'util';
import { ValidationError } from '../models/core';

@Injectable()
export class ErrorHandlerService {

  constructor(private snackBar: MatSnackBar,
              private router: Router,
              private session: SessionService) {}

  HandleError(error: Response) {
    /*
    console.log('ErrorHandler ErrorStatusText: ', error.statusText);
    console.log('ErrorHandler ErrorStatus: ', error.status);
    console.log('ErrorHandler ErrorType: ', error.type);
    console.log('ErrorHandler ErrorBody: ', error.body);
    console.log('Error: ', error);
    */

    // console.log('BODY = ', body);

    /*
    const conf = new MatSnackBarConfig();
    conf.duration = 100000;
    conf.panelClass = ['snackbar-success'];
    */
    let msg: string; msg = '';

    if ( error.status === 400 ) {
      /*
      let ves: ValidationError[] = x;
      console.log('x = ', ves);
      for ( let i = 0; i < ves.length; i++ ) {
        msg += `${ves[i].PropertyName}: ${ves[i].Error}` + '<br>';
      }
      */
     let ves: any; ves = error.json();
     for ( let i = 0; i < ves.length; i++ ) {
        msg += `${ves[i].PropertyName}: ${ves[i].Error}`;
     }
     // msg = '400 Bad Request';
     console.log('e = ', ves);

    } else {
      msg = `${error.status} ${error.json()}`;

      switch (error.status) {
        case 0:
          msg = '500 Server Internal Error';
          break;
        case 401:
          this.session.Close();
          this.router.navigate(['/login']);
          break;
        case 403:
          this.router.navigate(['/']);
          break;
      }
    }

    this.snackBar.openFromComponent(CustomSnackbarComponent, {
      panelClass: ['custom-snackbar-error'],
      data: {
        message: msg,
        icon: 'error',
        style: 'error'
      }
    });
  }
}
