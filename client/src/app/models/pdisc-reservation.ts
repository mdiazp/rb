export class PDiscReservation {
  constructor(
    public ID: number,
    public ClientID: number,
    public ClientName: string,
    public InitialTime: Date,
    public FinishTime: Date,
    public Cost: number,
    public TurnWeekDay: string,
    public TurnNum: number,
    public DiskCategoryRequest: string,
  ) {}
}

export class PDRFreeTurns {
  DisplayValue: string;

  constructor(
    public TurnWeekDay: string,
    public TurnNum: number,
  ) {
    this.DisplayValue = `${this.TurnWeekDay} - ${this.TurnNum}`;
  }
}
