import { AnyAction, DISPLAY_MESSAGE } from "./actions";
import { initialState, State } from "./state";

const reducer = (state: State = initialState, action: AnyAction) => {
  switch (action.type) {
    case DISPLAY_MESSAGE:
      return { ...state, message: action.message };
    default:
      return state;
  }
};

export default reducer;
