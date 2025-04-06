import { FaGithub, FaInstagram, FaPaypal, FaTwitter } from "react-icons/fa";
import "./Footer.css";

const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer-text">
        &copy; {new Date().getFullYear()}. AurisTFG - All Rights Reserved
      </div>
      <div className="social-icons">
        <a
          href="https://github.com/AurisTFG"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaGithub />
        </a>
        <a
          href="https://x.com/AurisTFG"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaTwitter />
        </a>
        <a
          href="https://www.instagram.com/auristfg"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaInstagram />
        </a>
        <a
          href="https://www.paypal.me/aurimasda"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FaPaypal />
        </a>
      </div>
    </footer>
  );
};

export default Footer;
