import axios, { AxiosResponse } from "axios";
import { FormField } from "../types/formField";

const api = axios.create({
  baseURL: process.env.PUBLIC_API_URL,
  headers: { "Content-Type": "application/json" },
});

async function login(
  payload: FormField
): Promise<{ username: string; password: string }> {
  try {
    const response: AxiosResponse<{ username: string; password: string }> =
      await api.post("/login", payload);
    return response.data; // Ensure this returns the expected data structure
  } catch (e) {
    console.error(e); // log for now
    throw e; // Re-throw the error for better handling in the hook
  }
}

const loginApis = {
  login,
};

export default loginApis;
