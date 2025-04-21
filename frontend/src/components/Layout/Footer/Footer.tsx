import { FaGithub, FaInstagram, FaPaypal, FaTwitter } from 'react-icons/fa';
import styles from './Footer.module.css';

export default function Footer() {
  return (
    <footer className={styles.footer}>
      <div className={styles.footerText}>
        &copy; {new Date().getFullYear()}. AurisTFG - All Rights Reserved
      </div>
      <div className={styles.socialIcons}>
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
}
