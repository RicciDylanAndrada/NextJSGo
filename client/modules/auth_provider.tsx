import React, { useState, createContext } from "react";
export type UserInfo = {
  username: string;
  id: string;
};

export const AuthContext = createContext<{
  authenticated: boolean;
  setAuthenticated: (auth: boolean) => void;
  user: UserInfo;
  setUser: (user: UserInfo) => void;
}>({
  authenticated: false,
  setAuthenticated: () => {},
  user: { username: "", id: "" },
  setUser: () => {},
});

function auth_provider({ children }: React.ReactNode) {
  const [authenticated, SetAuthenticated] = useState(false);
  const [user, setUser] = useState("");

  return <div>{children}</div>;
}

export default auth_provider;
