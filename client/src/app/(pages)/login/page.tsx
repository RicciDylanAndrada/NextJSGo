"use client";
import React, { useState } from "react";
import api from "../../api/login";
import { FormField } from "@/app/types/formField";
import { useRouter

 } from "next/router";
 import { UserInfo } from "../../../../modules/auth_provider";
 
export default function Page() {
  const router=useRouter()
  const [formField, setFormFields] = useState<FormField>({
    email: "",
    password: "",
  });

  const handleOnFormChange = (key: string, value: string) => {
    setFormFields(
      (prevFields) =>
        ({
          ...prevFields,
          [key]: value,
        } as FormField)
    ); //type assertion to override
  };
  async function submitHandler(e: React.SyntheticEvent) {
    e.preventDefault(); // so it does not relaod page

    try {
      if (formField?.email && formField?.password) {
        const response = await api.login(formField);
        const { username, password } = response;

        localStorage.setItem(
          "user_info",
          JSON.stringify({ username, password })
        );
        return router.push('/')
      }
    } catch (e) {
      alert(e);
    }
  }
  return (
    <div className="flex flex-col items-center justify-center min-w-full min-h-screen">
      <h1 className="text-3xl font-bold">Welcome</h1>
      <p className="text-sm">Please Login</p>
      <form className="w-full  flex flex-col md:w-1/5 ">
        <input
          placeholder="email"
          value={formField.email}
          onChange={(e) => {
            handleOnFormChange("email", e.target.value);
          }}
          className="rounded-md  bg-white mt-8 p-3 focus:outline-none focus:border-blue border-gray border-1"
        ></input>
        <input
          type="password"
          placeholder="password"
          value={formField.password}
          onChange={(e) => {
            handleOnFormChange("password", e.target.value);
          }}
          className="rounded-md bg-white  mt-8 p-3 focus:outline-none focus:border-blue border-gray border-1"
        ></input>
        <button
          type="submit"
          onClick={submitHandler}
          className="rounded-md  self-center w-1/2  mt-8 p-3 focus:outline-none focus:border-blue border-gray border-2"
        >
          login
        </button>
      </form>
    </div>
  );
}
