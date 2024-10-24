"use client";
import React, { useState } from "react";

type Props = {};

export default function Page({}: Props) {
  const [formField, setFormFields] = useState({
    email: "",
    password: "",
  });

  const handleOnFormChange = (key:String, value:String) => {
    setFormFields({ [key]: value });
  };
 async function  submitHandler(e:React.SyntheticEvent){
  e.preventDefault() // so it does not relaod page

  try{
const res= await fetch('',{

})
  }
  catch(e){
    alert(e)

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
        <button type='submit'
        onClick={submitHandler}
        className="rounded-md  self-center w-1/2  mt-8 p-3 focus:outline-none focus:border-blue border-gray border-2">
          login
        </button>
      </form>
    </div>
  );
}
