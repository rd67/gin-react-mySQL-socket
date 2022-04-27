export const toLocaleDTString = (
  value: any = Date.now(),
  format = "en-US",
  dateOptions: Intl.DateTimeFormatOptions = {
    dateStyle: "long",
  }
) => {
  let formatted = `${new Date(value).toLocaleDateString(
    format,
    dateOptions
  )}, ${new Date(value).toLocaleTimeString(format)}`;

  return formatted;
};
