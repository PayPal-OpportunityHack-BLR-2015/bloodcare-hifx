// Generated code from Butter Knife. Do not modify!
package com.hifx.user.bloodcare;

import android.view.View;
import butterknife.ButterKnife.Finder;
import butterknife.ButterKnife.Injector;

public class SignupActivity$$ViewInjector<T extends com.hifx.user.bloodcare.SignupActivity> implements Injector<T> {
  @Override public void inject(final Finder finder, final T target, Object source) {
    View view;
    view = finder.findRequiredView(source, 2131361901, "field '_nameText'");
    target._nameText = finder.castView(view, 2131361901, "field '_nameText'");
    view = finder.findRequiredView(source, 2131361895, "field '_emailText'");
    target._emailText = finder.castView(view, 2131361895, "field '_emailText'");
    view = finder.findRequiredView(source, 2131361896, "field '_passwordText'");
    target._passwordText = finder.castView(view, 2131361896, "field '_passwordText'");
    view = finder.findRequiredView(source, 2131361904, "field '_signupButton'");
    target._signupButton = finder.castView(view, 2131361904, "field '_signupButton'");
    view = finder.findRequiredView(source, 2131361905, "field '_loginLink'");
    target._loginLink = finder.castView(view, 2131361905, "field '_loginLink'");
    view = finder.findRequiredView(source, 2131361903, "field '_planets_spinner'");
    target._planets_spinner = finder.castView(view, 2131361903, "field '_planets_spinner'");
    view = finder.findRequiredView(source, 2131361902, "field '_input_phone'");
    target._input_phone = finder.castView(view, 2131361902, "field '_input_phone'");
  }

  @Override public void reset(T target) {
    target._nameText = null;
    target._emailText = null;
    target._passwordText = null;
    target._signupButton = null;
    target._loginLink = null;
    target._planets_spinner = null;
    target._input_phone = null;
  }
}
