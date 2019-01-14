import { Component, ViewChild, OnInit, AfterViewInit, ElementRef, AfterViewChecked, Input } from '@angular/core';
import { MatPaginator, MatSort, MatSelect, MatDatepicker, MatInput, MatDatepickerInputEvent } from '@angular/material';
import { tap, debounceTime, distinctUntilChanged } from 'rxjs/operators';
import { fromEvent, BehaviorSubject, Observable } from 'rxjs';
import { MatDialog } from '@angular/material';

import {
  APIPDiscReservationService,
  PDiscReservationFilter,
  PDiscReservationsDataSource,
  ErrorHandlerService,
  FeedbackHandlerService,
  Paginator,
  APIFreeInfoService,
  OrderBy,
} from '../../../services/core';
import { PDiscReservation, Util } from '../../../models/core';
import { isNullOrUndefined } from 'util';
import { Router, PreloadAllModules } from '@angular/router';
import { ClientSelectorComponent } from '../../my-common/client-selector/client-selector.component';

@Component({
  selector: 'app-pdisc-reservations-all',
  templateUrl: './pdisc-reservations-all.component.html',
  styleUrls: ['./pdisc-reservations-all.component.scss'],
})
export class PDiscReservationsAllComponent implements OnInit, AfterViewInit {

  dataSource: PDiscReservationsDataSource;
  displayedColumns= ['id', 'clientName', 'initialTime', 'finishTime', 'turn'];

  @ViewChild(MatPaginator) paginator: MatPaginator;

  /*
  public ActivedInitialTime: Date,
  public ActivedFinishTime: Date,
  */
  initialDateFilter: Date = null;
  finishDateFilter: Date = null;
  @ViewChild('weekDayFilter') weekDayFilter: MatSelect;
  @ViewChild('turnNumFilter') turnNumFilter: MatSelect;
  @ViewChild(ClientSelectorComponent) clientSelector: ClientSelectorComponent;
  @ViewChild('activedClientFilter') activedClientFilter: MatSelect;
  @ViewChild('discCategoryRequestFilter') discCategoryRequestFilter: MatSelect;

  completeFiltersInitialization = false;

  initialPageSize = 20;
  pageSizeOptions = [20, 50, 100];

  modelUtil = new Util();

  wds = this.modelUtil.GetWeekDaysInfo();
  turnNums = this.modelUtil.GetTurnNums();
  dcis = this.modelUtil.GetDiscCategoriesInfo();
  dcrn = this.modelUtil.GetDiskCategoryRequestNull();

  serverTime: Date = new Date();

  constructor(private api: APIPDiscReservationService,
              private free: APIFreeInfoService,
              private eh: ErrorHandlerService,
              private fh: FeedbackHandlerService,
              private router: Router,
              private dialog: MatDialog) {
    this.free.GetServerTime().subscribe(
      (date) => {
        this.serverTime = date;
      },
      (e) => {
        this.eh.HandleError(e);
      }
    );
  }

  ngOnInit() {
      this.dataSource = new PDiscReservationsDataSource(this.api, this.eh);
      this.dataSource.load(
        true,
        new PDiscReservationFilter(
          null, true, null, null, null, null, null,
          new Paginator(
            0,
            this.initialPageSize,
          ),
          new OrderBy('initial_time', false),
        ),
      );
  }

  ngAfterViewInit(): void {
    this.weekDayFilter.valueChange.subscribe(() => this.ToFilter());
    this.turnNumFilter.valueChange.subscribe(() => this.ToFilter());
    this.clientSelector.selectionChanges.subscribe(() => this.ToFilter());
    this.activedClientFilter.valueChange.subscribe(() => this.ToFilter());
    this.discCategoryRequestFilter.valueChange.subscribe(() => this.ToFilter());
    this.paginator.page.pipe(tap(() => this.load(false))).subscribe();

    this.completeFiltersInitialization = true;
  }

  ToFilter(): void {
    this.paginator.pageIndex = 0;
    this.load(true);
  }

  InitialDatePickerChange(ev: MatDatepickerInputEvent<Date>): void {
    this.initialDateFilter = ev.value;
    this.ToFilter();
  }

  FinishDatePickerChange(ev: MatDatepickerInputEvent<Date>): void {
    this.finishDateFilter = ev.value;
    this.ToFilter();
  }

  load(loadCount: boolean) {
    if ( !this.completeFiltersInitialization ) {
      return;
    }

    let clientID: number; clientID = 0;
    if ( this.clientSelector.ValidSelection() ) {
      clientID = this.clientSelector.autoClientSelection.ID;
    }

    this.dataSource.load(
      loadCount,
      new PDiscReservationFilter(
        clientID,
        (this.activedClientFilter.value === 'all' ? null : this.activedClientFilter.value),
        this.initialDateFilter,
        this.finishDateFilter,
        (this.discCategoryRequestFilter.value === 'all' ? null : this.discCategoryRequestFilter.value),
        (this.weekDayFilter.value === 'all' ? null : this.weekDayFilter.value),
        (this.turnNumFilter.value === 'all' ? null : this.turnNumFilter.value),
        new Paginator(
          this.paginator.pageIndex * this.paginator.pageSize,
          this.paginator.pageSize
        ),
        new OrderBy('initial_time', false),
      ),
    );
  }

  DateFormat(d: Date): string {
    return this.modelUtil.FormatDateToDisplay(d);
  }
}
