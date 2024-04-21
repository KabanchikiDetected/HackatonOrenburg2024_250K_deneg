import "./index.scss";

const Modal = ({ children, isOpen, setIsOpen }: any) => {
  return (
    <div className={isOpen ? "modal active" : "modal"} onClick={() => setIsOpen(false)}>
      <div className="modal__content"
        onClick={(event) => event.stopPropagation()}
      >{children}</div>
    </div>
  );
};

export default Modal;
