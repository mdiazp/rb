export interface ValueAndDisplay {
  Value: string;
  DisplayValue: string;
}

export interface WeekDayInfo {
  Value: string;
  DisplayValue: string;
}

export class Util {
  constructor() {}

  GetDiscCategoriesInfo(): ValueAndDisplay[] {
    return [
      { Value: 'BIG', DisplayValue: 'Grande' },
      { Value: 'SMALL', DisplayValue: 'Chiquito' },
    ];
  }

  GetDiskCategoryRequestNull(): ValueAndDisplay {
    return { Value: 'NULL', DisplayValue: 'Sin especificar' };
  }

  GetTurnNums(): number[] {
    return [1, 2, 3, 4, 5];
  }

  GetWeekDaysInfo(): ValueAndDisplay[] {
    return [
      { Value: 'Monday', DisplayValue: 'Lunes' },
      { Value: 'Tuesday', DisplayValue: 'Martes' },
      { Value: 'Wednesday', DisplayValue: 'Miercoles' },
      { Value: 'Thursday', DisplayValue: 'Jueves' },
      { Value: 'Friday', DisplayValue: 'Viernes' },
      { Value: 'Saturday', DisplayValue: 'Sabado' },
      { Value: 'Sunday', DisplayValue: 'Domingo' },
    ];
  }

  GetWeekDayDisplayValue(weekdayValue: string): string {
    const wds = this.GetWeekDaysInfo();
    for ( let i = 0; i < wds.length; i++ ) {
      if ( wds[i].Value === weekdayValue ) {
        return wds[i].DisplayValue;
      }
    }
    return weekdayValue;
  }

  FormatDateToSendToApi(date: Date): string {
    // 2019-01-09
    return date.getFullYear() + '-' + this.dd(date.getMonth() + 1) + '-' +
           this.dd(date.getDate());
  }

  FormatDateToDisplay(date: Date): string {
    return  this.dd(date.getDate())  + '/' + this.dd(date.getMonth() + 1) + '/' +
            date.getFullYear();
  }

  NewDate(s: string): Date {
    // 2000-01-01
    // 0123456789
    s = s.substring(0, 10);
    const xs = s.split('-');
    return new Date(Number(xs[0]), Number(xs[1]) - 1, Number(xs[2]));
  }

  private dd(x: number): string {
    if (x >= 0 && x < 10) {
      return '0' + x.toString();
    }
    return x.toString();
  }
}
