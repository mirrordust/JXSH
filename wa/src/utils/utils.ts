import {
  addHours,
  formatDistanceToNow,
  parseISO
} from 'date-fns';


/**
  * Correct the time because the backend
  * UTC time is parsed as with timezone 
  */
export function correctDate(timestamp: string): string;
export function correctDate(timestamp: null): null;
export function correctDate(timestamp: string | null): string | null {
  if (timestamp) {
    const date = parseISO(timestamp);
    const correctedDate = addHours(date, -date.getTimezoneOffset() / 60);
    return correctedDate.toLocaleDateString()
      + ' ' + correctedDate.toLocaleTimeString();
  }
  return null;
}

export function timeAgo(timestamp: string | null): string {
  let timeAgo = '';
  if (timestamp) {
    const date = parseISO(timestamp);
    const correctedDate = addHours(date, -date.getTimezoneOffset() / 60);
    const timePeriod = formatDistanceToNow(correctedDate);
    timeAgo = `${timePeriod} ago`;
  }
  return timeAgo;
};
