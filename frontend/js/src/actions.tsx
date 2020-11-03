import { Action } from "redux";

export const DISPLAY_MESSAGE = "DISPLAY_MESSAGE";

export interface DisplayMessageAction extends Action {
  type: "DISPLAY_MESSAGE";
  message: string;
}

export const displayMessage = (message: string) => ({
  type: DISPLAY_MESSAGE,
  message,
});

// Export union type
export type AnyAction = DisplayMessageAction;
