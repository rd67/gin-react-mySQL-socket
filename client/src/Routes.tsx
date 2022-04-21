import { Routes, Route } from "react-router-dom";

import { NotFoundComp } from "./components";

import Home from "./pages/Home/Home";

export default function () {
  return (
    <Routes>
      <Route path="/" element={<Home />} />

      <Route path="*" element={<NotFoundComp />} />
    </Routes>
  );
}
