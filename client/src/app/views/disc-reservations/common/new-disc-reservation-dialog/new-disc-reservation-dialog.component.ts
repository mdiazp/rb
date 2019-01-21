import { Component, OnInit, ViewChild, Inject, AfterViewInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ClientSelectorComponent } from '../../../my-common/client-selector/client-selector.component';
import { Util } from '../../../../models/util';
import { PDRFreeTurns, PDiscReservation } from '../../../../models/core';
import { APIPDiscReservationService, ErrorHandlerService, FeedbackHandlerService, APIFreeInfoService } from '../../../../services/core';
import { Router } from '@angular/router';
import { MAT_DIALOG_DATA, MatDialog, MatDialogRef } from '@angular/material';
import { isNullOrUndefined } from 'util';

@Component({
  selector: 'app-new-disc-reservation-dialog',
  templateUrl: './new-disc-reservation-dialog.component.html',
  styleUrls: ['./new-disc-reservation-dialog.component.css']
})
export class NewDiscReservationDialogComponent implements OnInit, AfterViewInit {

  form: FormGroup;
  @ViewChild('app-client-selector') clientSelector: ClientSelectorComponent;
  initialTime: FormControl;
  finishTime: FormControl;
  cost: FormControl;
  dcr: FormControl;
  ft: FormControl;

  serverTime: Date = new Date();

  modelUtil = new Util();
  freeTurns: PDRFreeTurns[] = [];
  routeToGetBack: string[] = [];

  discCategories = this.modelUtil.GetDiscCategoriesInfo();

  constructor(private api: APIPDiscReservationService,
              private freeAPI: APIFreeInfoService,
              private eh: ErrorHandlerService,
              private feedback: FeedbackHandlerService,
              private router: Router,
              private ref: MatDialogRef<NewDiscReservationDialogComponent>,
              @Inject(MAT_DIALOG_DATA) public data: any) {
    this.routeToGetBack = data.routeToGetBack;
  }

  ngOnInit() {
    this.initForm();
    this.loadServerTime();
    // this.loadFreeTurns();
    this.feedback.ShowFeedback('dfnvidfjvfbjbhjb');
  }

  ngAfterViewInit(): void {}

  initForm(): void {
    this.initialTime = new FormControl(this.serverTime, Validators.required);
    this.finishTime = new FormControl();
    this.cost = new FormControl(20, Validators.required);
    this.dcr = new FormControl(
      this.modelUtil.GetDiskCategoryRequestNull().Value,
      Validators.required
    );
    this.ft = new FormControl('', Validators.required);

    this.form = new FormGroup(
      {
        'initialTime' : this.initialTime,
        'finishTime' : this.finishTime,
        'cost' : this.cost,
        'dcr': this.dcr,
        'ft' : this.ft,
      }
    );
  }

  EnableSubmitButton(): boolean {
    return (
      (!isNullOrUndefined(this.form) && this.form.valid) &&
      (!isNullOrUndefined(this.clientSelector) && this.clientSelector.ValidSelection())
    );
  }

  onSubmit(): void {
    const pdr = new PDiscReservation(
      0,
      this.clientSelector.autoClientSelection.ID,
      '',
      new Date(this.initialTime.value),
      new Date(this.finishTime.value),
      Number(this.cost.value),
      this.ft.value.TurnWeekDay,
      this.ft.value.TurnNum,
      this.dcr.value,
    );
    this.api.PostPDiscReservation(pdr).subscribe(
      (_) => {
        this.feedback.ShowFeedback('La nueva reservacion fue creada exitosamente.');
        this.router.navigate(this.routeToGetBack);
        this.ref.close();
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  loadServerTime(): void {
    this.freeAPI.GetServerTime().subscribe(
      (st) => {
        this.serverTime = st;
        this.initialTime.setValue(this.serverTime);
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  loadFreeTurns(): void {
    this.api.GetFreeTurnsToAdd(
      this.initialTime.value,
      this.finishTime.value,
      this.dcr.value,
    ).subscribe(
      (freeTurns) => {
        this.freeTurns = freeTurns;
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }
}
