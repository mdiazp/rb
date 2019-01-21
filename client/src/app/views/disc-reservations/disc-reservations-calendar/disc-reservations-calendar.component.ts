import { Component, OnInit, ViewChild, AfterViewChecked, AfterViewInit } from '@angular/core';
import { Util, ValueAndDisplay } from '../../../models/util';
import { MatSelect } from '@angular/material';
import { ErrorHandlerService, APIPDiscReservationService, APIFreeInfoService } from '../../../services/core';
import { BehaviorSubject } from 'rxjs';
import { PDRTurnCalendar, PDRTurnCalendarState } from '../../../models/core';
import { isNullOrUndefined } from 'util';

@Component({
  selector: 'app-disc-reservations-calendar',
  templateUrl: './disc-reservations-calendar.component.html',
  styleUrls: ['./disc-reservations-calendar.component.css']
})
export class DiscReservationsCalendarComponent implements OnInit, AfterViewInit {

  private loadingSubject = new BehaviorSubject<boolean>(true);
  loading$ = this.loadingSubject.asObservable();

  modelUtil = new Util();
  wds = this.modelUtil.GetWeekDaysInfo();
  initialWD = this.wds[0];
  turnNums = this.modelUtil.GetTurnNums();
  initialTN = this.turnNums[0];

  currentDate = new Date();

  calendar = new PDRTurnCalendar(0, []);
  filteredStates: PDRTurnCalendarState[] = [];

  @ViewChild('weekDayFilter') weekDayFilter: MatSelect;
  @ViewChild('turnNumFilter') turnNumFilter: MatSelect;

  statusOptionAll = true;
  statusOptionWrong = false;

  constructor(private api: APIPDiscReservationService,
              private apiFree: APIFreeInfoService,
              private eh: ErrorHandlerService) {}

  ngOnInit() {}

  ngAfterViewInit(): void {
    this.apiFree.GetServerTime().subscribe(
      (currentDate) => {
        this.currentDate = currentDate;
        this.initialWD = this.modelUtil.GetWeekDaysInfo()[this.currentDate.getDay()];

        this.weekDayFilter.valueChange.subscribe(() => this.load());
        this.turnNumFilter.valueChange.subscribe(() => this.load());
        console.log('wd = ', this.initialWD);
        this.load();
      },
      (e) => {
        this.weekDayFilter.valueChange.subscribe(() => this.load());
        this.turnNumFilter.valueChange.subscribe(() => this.load());
        this.eh.HandleError(e);
      }
    );
  }

  statusFilterOnSelect(option: string): void {
    if ((option === 'all' && this.statusOptionAll) ||
        (option !== 'all' && !this.statusOptionAll)) {
      return;
    }
    this.statusOptionAll = !this.statusOptionAll;
    this.statusOptionWrong = !this.statusOptionWrong;

    this.filter();
  }

  filter(): void {
    if ( this.statusOptionAll ) {
      this.filteredStates = this.calendar.States;
      return;
    }
    let tmp: PDRTurnCalendarState[]; tmp = [];
    for ( let i = 0; i < this.calendar.States.length; i++ ) {
      if ( this.calendar.States[i].Wrong ) {
        tmp.push(this.calendar.States[i]);
      }
    }
    this.filteredStates = tmp;
  }

  load(): void {
    console.log('load');
    this.loadingSubject.next(true);
    this.api.GetCalendar(
      this.initialWD.Value,
      this.initialTN,
      // (isNullOrUndefined(this.weekDayFilter) ? this.initialWD.Value : this.weekDayFilter.value),
      // (isNullOrUndefined(this.turnNumFilter) ? this.initialTN : this.turnNumFilter.value),
    )
    .subscribe(
      (calendar) => {
        this.calendar = calendar;
        this.filter();
        this.loadingSubject.next(false);
      },
      (e) => {
        this.eh.HandleError(e);
        this.loadingSubject.next(false);
      }
    );
  }

  isGoodState(state: PDRTurnCalendarState): boolean {
    if ( state.DiscsTotal < state.PDRs.length ) {
      return false;
    }

    for ( let i = 0; i < state.DCRR.length; i++ ) {
      if ( state.DCRR[i].DCTotal < state.DCRR[i].DCRTotal ) {
        return false;
      }
    }
    return true;
  }
}
