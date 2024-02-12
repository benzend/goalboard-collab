import { ComponentProps, FunctionComponent } from 'react';
import cn from 'classnames';

export interface HeadingProps extends ComponentProps<HeadingType> {
  el: HeadingType;
  type: HeadingType;
}

export const Heading: FunctionComponent<HeadingProps> = (props) => {
  const extended: ComponentProps<HeadingType> = {
    ...props,
    className: cn(getHeadingClasses(props.type), props.className),
  };
  return <props.el {...extended}>{props.children}</props.el>;
};

export function getHeadingClasses(style: HeadingType): string {
  switch (style) {
    case 'h1':
      return 'text-4xl font-bold';
    case 'h2':
      return 'text-xl font-semibold';
    case 'h3':
      return 'text-lg font-semibold';
    case 'h4':
    default:
      return 'text-md font-semibold';
  }
}

export type HeadingType = 'h1' | 'h2' | 'h3' | 'h4' | 'h5' | 'h6';
