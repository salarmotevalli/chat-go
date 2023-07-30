import React from "react";
import style from "./Form.module.css";
import logo from "../assets/logo.svg";
import RegisterForm from "./RegisterForm";

const RegisterFormContainer = () => {
  return (
    <>
      <div className="row justify-content-center mt-3">
        <div
          className={`col-11 col-sm-10 col-md-8 col-lg-6 col-xl-5 rounded-5 mt-5 ${style.form_bg}`}
        >
          <div className="brand d-flex justify-content-center align-items-center flex-column flex-sm-row mt-0 mb-2 mt-sm-5 mb-sm-4">
            <img src={logo} alt="logo" className={` ${style.logo}`} />
            <h1 className="text-uppercase text-wrap text-white ms-3 mt-3 mt-sm-2">
              snappy
            </h1>
          </div>
          <div className="container">
            <RegisterForm />
          </div>
        </div>
      </div>{" "}
    </>
  );
};

export default RegisterFormContainer;
