import { PDiscReservation } from './pdisc-reservation';

export class PDRTurnCalendarState {
  constructor(
    public Date: Date,
    public PDRs: PDiscReservation[],
    public DCRR: DiscCategoryRequestReport[],
    public DCRNullTotal: number,
    public DiscsTotal: number) {}
}

export class DiscCategoryRequestReport {
  constructor(
    public Category: string,
    public DCTotal: number,
    public DCRTotal: number,
  ) {}
}
