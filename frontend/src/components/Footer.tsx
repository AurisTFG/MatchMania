import "./Footer.css";

const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer-text">
        &copy; {new Date().getFullYear()}. IFF-1/1 gr. stud. Aurimas Dabrišius
      </div>
    </footer>
  );
};

export default Footer;
