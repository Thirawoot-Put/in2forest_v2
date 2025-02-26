"use client"

import InputLogin from "@/components/input/loginInput";
import { Button, Divider } from "@mui/material";
import React, { useState } from "react";

type TLoginInput = {
  username: string,
  password: string
}

export default function FormLoginAdmin() {
  const [input, setInput] = useState<TLoginInput>({ username: "", password: "" })

  const onClickLogin = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    console.log(input)
  }

  const onClickRegister = () => {
    console.log('click register!!!')
  }

  const handleChangeInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setInput({ ...input, [name]: value })
  }

  return (
    <form className="w-full h-full py-4" onSubmit={onClickLogin}>
      <div className="flex flex-col gap-2 items-center ">
        <InputLogin name="username" label="Username" size="small" placeholder="sample@mail.com" type="" onChange={handleChangeInput} />
        <InputLogin name="password" label="Password" size="small" placeholder="password1234" type="password" onChange={handleChangeInput} />
      </div>
      <div className="flex flex-col items-center justify-center gap-2 pt-5">
        <Button type="submit" variant="contained">Login</Button>
        <Divider className="w-3/5">or</Divider>
        <Button type="button" onClick={onClickRegister} variant="text">Register</Button>
      </div>
    </form >
  )
}
