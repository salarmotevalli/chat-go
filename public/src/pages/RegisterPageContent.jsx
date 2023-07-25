import React from "react";
import logo from "../assets/logo.svg";
import style from "./Login.module.css";
import { Link } from "react-router-dom";

const RegisterPageContent = ({ handleChange, handleSubmit }) => {
  return (
    <>
      <div className="container">
        <div className="row vh-100 justify-content-center align-items-center">
          <div
            className={`col-10 col-sm-9 col-md-8 col-lg-6 col-xl-5 rounded-5 ${style.form_bg}`}
          >
            <div className="row justify-content-center align-items-center">
              <div className="col-12 mt-4">
                <div className="brand d-flex justify-content-center align-items-center">
                  <img src={logo} alt="logo" className={style.logo} />
                  <h1 className="text-uppercase text-wrap text-white ms-3">
                    snappy
                  </h1>
                </div>
              </div>
            </div>
            <div className="row justify-content-center">
              <div className="col-10 col-sm-10 col-md-9 p-0 my-4">
                <form
                  action=""
                  onSubmit={(event) => handleSubmit(event)}
                  className="form-control border-0 d-flex flex-column bg-transparent"
                >
                  <input
                    className={`bg-transparent rounded mb-4 text-white ${style.input}`}
                    type="text"
                    placeholder="Username"
                    name="username"
                    onChange={(e) => handleChange(e)}
                  />
                  <input
                    className={`bg-transparent rounded mb-4 text-white ${style.input}`}
                    type="email"
                    placeholder="Email"
                    name="email"
                    onChange={(e) => handleChange(e)}
                  />
                  <input
                    className={`bg-transparent rounded mb-4 text-white ${style.input}`}
                    type="password"
                    placeholder="Password"
                    name="password"
                    onChange={(e) => handleChange(e)}
                  />
                  <input
                    className={`bg-transparent rounded mb-4 text-white ${style.input}`}
                    type="password"
                    placeholder="Confirm Password"
                    name="confirmPassword"
                    onChange={(e) => handleChange(e)}
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
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default RegisterPageContent;
