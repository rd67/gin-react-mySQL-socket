import { toast, ToastOptions } from 'react-toastify';

export const successToast = (message: string, options?: ToastOptions) =>
  toast.success(message, options);

export const errorToast = (message: string, options?: ToastOptions) =>
  toast.error(message, options);

export const warningToast = (message: string, options?: ToastOptions) =>
  toast.warning(message, options);

export const infoToast = (message: string, options?: ToastOptions) =>
  toast.info(message, options);

export const clearAllToast = () => toast.dismiss();
