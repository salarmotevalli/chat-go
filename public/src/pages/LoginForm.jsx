import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import style from "./Form.module.css";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";
import { loginRoute } from "../utils/APIRoutes";
import { useNavigate } from "react-router-dom";

const schema = z.object({
  username: z
    .string()
    .min(4, { message: "The username must be 4 characters or more" })
    .max(10, { message: "The username must be 10 characters or less" })
    .regex(
      /^[a-zA-Z0-9_]+$/,
      "The username must contain only letters, numbers and underscore (_)"
    ),
  password: z
    .string()
    .min(1, "Password is required")
    .min(8, "Password must have more than 8 characters"),
});

const LoginForm = () => {
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
      const { username, password } = values;
      const { data: response } = await axios.post(loginRoute, {
        username,
        password,
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
              {...register("password")}
              type="password"
              placeholder="Password"
              name="password"
              className={`w-100 rounded bg-transparent my-2 my-sm-3 text-white ${style.input}`}
            />
            {errors.password && (
              <p className="text-danger">{errors.password.message}</p>
            )}
            <button
              type="submit"
              className={`border-0 rounded text-uppercase fw-bold text-white my-3 my-sm-4 ${style.login_btn}`}
            >
              Login
            </button>
            <span
              className={`mb-5 text-center text-white text-uppercase ${style.link_text}`}
            >
              Don't have an account ?{" "}
              <Link
                to="/register"
                className={`fw-bold text-decoration-none ${style.link}`}
              >
                Create One.
              </Link>
            </span>{" "}
          </form>
        </div>
      </div>
    </>
  );
};

export default LoginForm;
