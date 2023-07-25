import React from "react";
import style from "./Login.module.css";
import { Link } from "react-router-dom";

const RegisterForm = () => {
  return (
    <>
      <form
        action=""
        className="form-control border-0 d-flex flex-column bg-transparent"
      >
        <input
          className={`bg-transparent rounded mb-4 text-white ${style.input}`}
          type="text"
          placeholder="Username"
          name="username"
        />
        <input
          className={`bg-transparent rounded mb-4 text-white ${style.input}`}
          type="email"
          placeholder="Email"
          name="email"
        />
        <input
          className={`bg-transparent rounded mb-4 text-white ${style.input}`}
          type="password"
          placeholder="Password"
          name="password"
        />
        <input
          className={`bg-transparent rounded mb-4 text-white ${style.input}`}
          type="password"
          placeholder="Confirm Password"
          name="confirmPassword"
        />
        <button
          type="submit"
          className={`rounded mb-4 text-white text-uppercase fw-bold border-0 ${style.login_btn}`}
        >
          Create User
        </button>
        <span className=" text-white text-uppercase">
          Don't have an account ?{" "}
          <Link
            to="/login"
            className={`fw-bold text-decoration-none ${style.link}`}
          >
            Login.
          </Link>
        </span>{" "}
      </form>
    </>
  );
};

export default RegisterForm;
