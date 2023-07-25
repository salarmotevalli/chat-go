import React from "react";
import logo from "../assets/logo.svg";
import style from "./Login.module.css";
import RegisterForm from "./RegisterForm";

const RegisterFormContainer = () => {
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
                <RegisterForm />
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default RegisterFormContainer;
