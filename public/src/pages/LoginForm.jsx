import React from "react";
import { Link } from "react-router-dom";
import style from "./Form.module.css";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import useCustomLogin from "../hooks/useCustomLogin";

const LoginForm = () => {
  const { schema, errorMessage, onSubmit } = useCustomLogin();

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
                to="/Register"
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
