import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";
import { registerRoute } from "../utils/APIRoutes";
import style from "./Form.module.css";
import axios from "axios";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

const schema = z
  .object({
    username: z
      .string()
      .min(4, { message: "The username must be 4 characters or more" })
      .max(10, { message: "The username must be 10 characters or less" })
      .regex(/^[a-zA-Z0-9_]+$/, {
        message:
          "The username must contain only letters, numbers and underscore (_)",
      }),
    email: z
      .string()
      .email({ message: "Invalid email" })
      .min(1, { message: "Email is required" }),
    password: z
      .string()
      .min(1, { message: "Password is required" })
      .min(8, { message: "Password must have more than 8 characters" }),
    confirmPassword: z
      .string()
      .min(1, { message: "Password confirmation is required" }),
  })
  .refine((data) => data.password === data.confirmPassword, {
    path: ["confirmPassword"],
    message: "Passwords do not match",
  });

const RegisterForm = () => {
  const [errorMessage, setErrorMessage] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    if (localStorage.getItem(process.env.REACT_APP_LOCALHOST_KEY)) {
      navigate("/");
    }
  }, []);

  const onSubmit = async (data) => {
    try {
      const values = schema.parse(data);
      const { username, email, password, confirmPassword } = values;
      const { data: response } = await axios.post(registerRoute, {
        username,
        email,
        password,
        confirmPassword,
      });
      if (response.status === false) {
        setErrorMessage(response.msg);
      }
      if (response.status === true) {
        localStorage.setItem(
          process.env.REACT_APP_LOCALHOST_KEY,
          JSON.stringify(response.user)
        );
        navigate("/");
      }
    } catch (error) {
      console.error(error);
    }
  };
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({ resolver: zodResolver(schema) });
  return (
    <>
      <div className="row justify-content-center">
        <div className="col-12 col-md-10">
          <form
            onSubmit={handleSubmit(onSubmit)}
            className="form-control bg-transparent d-flex justify-content-center flex-column border-0"
          >
            {errorMessage && <p className="text-danger">{errorMessage}</p>}
            <input
              {...register("username")}
              type="text"
              placeholder="Username"
              name="username"
              className={`w-100 rounded bg-transparent mb-3 text-white ${style.input}`}
            />
            {errors.username && (
              <p className="text-danger">{errors.username.message}</p>
            )}
            <input
              {...register("email")}
              type="email"
              placeholder="example@gmail.com"
              name="email"
              className={`w-100 rounded bg-transparent my-2 my-sm-3 text-white ${style.input}`}
            />
            {errors.email && (
              <p className="text-danger">{errors.email.message}</p>
            )}
            <input
              {...register("password")}
              type="password"
              placeholder="Password"
              name="password"
              className={`w-100 rounded bg-transparent my-2 my-sm-3 text-white ${style.input}`}
            />
            {errors.password && (
              <p className="text-danger">{errors.password.message}</p>
            )}
            <input
              {...register("confirmPassword")}
              type="password"
              placeholder="confirmPassword"
              name="confirmPassword"
              className={`w-100 rounded bg-transparent my-2 my-sm-3 text-white ${style.input}`}
            />
            {errors.confirmPassword && (
              <p className="text-danger">{errors.confirmPassword.message}</p>
            )}
            <button
              type="submit"
              className={`border-0 rounded text-uppercase fw-bold text-white my-3 my-sm-4 ${style.login_btn}`}
            >
              create user
            </button>
            <span
              className={`mb-5 text-center text-white text-uppercase ${style.link_text}`}
            >
              already have an account ?{" "}
              <Link
                to="/Login"
                className={`fw-bold text-decoration-none ${style.link}`}
              >
                login.
              </Link>
            </span>{" "}
          </form>
        </div>
      </div>
    </>
  );
};

export default RegisterForm;
