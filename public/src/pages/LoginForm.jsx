import React from "react";
import { Link } from "react-router-dom";
import style from "./Form.module.css";

const LoginForm = () => {
  return (
    <>
      <div className="row justify-content-center">
        <div className="col-12 col-md-10">
          <form className="form-control bg-transparent d-flex justify-content-center flex-column border-0">
            <input
              type="text"
              placeholder="Username"
              name="username"
              className={`w-100 rounded bg-transparent mb-3 text-white ${style.input}`}
            />
            <input
              type="password"
              placeholder="Password"
              name="password"
              className={`w-100 rounded bg-transparent my-2 my-sm-3 text-white ${style.input}`}
            />
            <button
              type="submit"
              className={`border-0 rounded text-uppercase fw-bold text-white my-3 my-sm-4 ${style.login_btn}`}
            >
              Login
            </button>
            <span
              className={`mb-4 text-center text-white text-uppercase ${style.link_text}`}
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
