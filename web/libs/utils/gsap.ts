import { gsap } from 'gsap';
import { ScrollTrigger } from 'gsap/dist/ScrollTrigger';
import { ScrollToPlugin } from 'gsap/dist/ScrollToPlugin';
import { CustomEase } from 'gsap/dist/CustomEase';

gsap.registerPlugin(ScrollTrigger, CustomEase);
gsap.registerPlugin(ScrollTrigger, ScrollToPlugin);

export { ScrollTrigger, CustomEase };

export default gsap;
