import './index.scss';

const Footer = () => {
  return (
    <footer className="footer">
      <div className="footer__logo">
        <img src="/images/small_logo.png" alt="" />
      </div>
      <div className="policy">Политика конфиденциальности</div>
      <div className="footer__group-links">
        <div className="group-links__item">О нас</div>
        <div className="group-links__item">Реклама</div>
        <div className="group-links__item">Партнеры</div>
      </div>
    </footer>
  );
};

export default Footer;