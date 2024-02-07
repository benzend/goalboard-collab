import { Link } from 'react-router-dom';

export const Nav = () => {
  return (
    <nav className="bg-cyan-600 p-4 flex justify-between align-middle">
      <section className="flex align-middle">
        <Link to="/">
          <span className="text-lg font-bold text-white">Goalboard</span>
        </Link>
      </section>
    </nav>
  );
};
